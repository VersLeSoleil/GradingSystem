from flask import Flask, request, jsonify
from PIL import Image
import torch
import torch.nn as nn
from torchvision import transforms
from efficientnet_pytorch import EfficientNet

app = Flask(__name__)

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")

model = EfficientNet.from_name('efficientnet-b0')
model._fc = nn.Linear(model._fc.in_features, 2)
model.load_state_dict(torch.load('best_model.pth', map_location=device))
model = model.to(device)
model.eval()

transform = transforms.Compose([
    transforms.Resize((224, 224)),
    transforms.ToTensor(),
    transforms.Lambda(lambda x: x.repeat(3, 1, 1)),
    transforms.Normalize([0.5]*3, [0.5]*3)
])

mask_transform = transforms.Compose([
    transforms.Resize((224, 224)),
    transforms.ToTensor()
])

class_names = ['良性', '恶性']

@app.route("/predict", methods=["POST"])
def predict():
    images = request.files.getlist('images')
    masks = request.files.getlist('masks')

    if not images or not masks or len(images) != len(masks):
        return jsonify({"error": "缺少 images 或 masks 文件，或数量不一致"}), 400

    results = []
    for img_file, mask_file in zip(images, masks):
        try:
            image = Image.open(img_file.stream).convert("L")
            mask = Image.open(mask_file.stream).convert("L")

            image_tensor = transform(image)
            mask_tensor = mask_transform(mask)
            mask_tensor = (mask_tensor > 0).float()
            image_tensor = image_tensor * mask_tensor.repeat(3, 1, 1)

            image_tensor = image_tensor.unsqueeze(0).to(device)

            with torch.no_grad():
                outputs = model(image_tensor)
                pred = torch.argmax(outputs, dim=1).item()
                confidence = torch.softmax(outputs, dim=1)[0, pred].item()

            results.append({
                "filename": img_file.filename,
                "class": class_names[pred],
                "confidence": round(confidence, 4)
            })
        except Exception as e:
            results.append({
                "filename": img_file.filename,
                "error": str(e)
            })

    return jsonify({"results": results})

if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
