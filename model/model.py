from flask import Flask, request, jsonify
from PIL import Image
import torch
import torch.nn as nn
from torchvision import transforms
from efficientnet_pytorch import EfficientNet
import io
import os

# 创建 Flask 应用
app = Flask(__name__)

# 设备设置
device = torch.device("cuda" if torch.cuda.is_available() else "cpu")

# 模型加载
model = EfficientNet.from_name('efficientnet-b0')
model._fc = nn.Linear(model._fc.in_features, 2)  # 二分类
model.load_state_dict(torch.load('best_model.pth', map_location=device))
model = model.to(device)
model.eval()

# 图像预处理
transform = transforms.Compose([
    transforms.Resize((224, 224)),
    transforms.ToTensor(),
    transforms.Lambda(lambda x: x.repeat(3, 1, 1)),  # 灰度转3通道
    transforms.Normalize([0.5] * 3, [0.5] * 3)
])

# 掩膜预处理（可选，假设上传时也包含 mask）
mask_transform = transforms.Compose([
    transforms.Resize((224, 224)),
    transforms.ToTensor()
])

# 类别映射
class_names = ['良性', '恶性']


@app.route("/predict", methods=["POST"])
def predict():
    if 'image' not in request.files or 'mask' not in request.files:
        return jsonify({"error": "缺少 image 或 mask 文件"}), 400

    image_file = request.files['image']
    mask_file = request.files['mask']

    try:
        # 读取图像与掩膜
        image = Image.open(image_file.stream).convert("L")
        mask = Image.open(mask_file.stream).convert("L")

        # 预处理
        image_tensor = transform(image)       # [3, 224, 224]
        mask_tensor = mask_transform(mask)    # [1, 224, 224]
        mask_tensor = (mask_tensor > 0).float()
        image_tensor = image_tensor * mask_tensor.repeat(3, 1, 1)

        image_tensor = image_tensor.unsqueeze(0).to(device)  # 添加 batch 维度

        # 模型推理
        with torch.no_grad():
            outputs = model(image_tensor)
            pred = torch.argmax(outputs, dim=1).item()
            confidence = torch.softmax(outputs, dim=1)[0, pred].item()

        return jsonify({
            "class": class_names[pred],
            "confidence": round(confidence, 4)
        })

    except Exception as e:
        return jsonify({"error": str(e)}), 500


if __name__ == "__main__":
    app.run(host="0.0.0.0", port=5000)
