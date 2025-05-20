from radiomics import featureextractor
import os
import SimpleITK as sitk
import pandas as pd


# 设置数据路径
basePath = r'E:/dachuang/Radiomics/Data/TestData'
imageFile = os.path.join(basePath, 'AP1.jpg')  # 原始图像
maskFile = os.path.join(basePath, 'AP1.nii.gz')  # 原始掩膜

# 读取图像和掩膜
image = sitk.ReadImage(imageFile)
mask = sitk.ReadImage(maskFile)

# 检查图像维度和类型
print("Original Image size:", image.GetSize())
print("Original Image pixel type:", image.GetPixelIDTypeAsString())
print("Mask size:", mask.GetSize())

# 如果图像不是标量图像，转换为灰度图像
if image.GetNumberOfComponentsPerPixel() > 1:
    print("Converting image to grayscale...")
    image = sitk.VectorIndexSelectionCast(image, 0)  # 提取第一个通道
    convertedImagePath = os.path.join(basePath, 'AP1_gray.nii.gz')
    sitk.WriteImage(image, convertedImagePath)  # 保存为灰度图像
    imageFile = convertedImagePath  # 更新图像路径

# 如果掩膜是三维的，转换为二维
if len(mask.GetSize()) == 3 and mask.GetSize()[-1] == 1:
    print("Converting mask to 2D...")
    mask = sitk.Extract(mask, mask.GetSize()[:-1] + (0,))
    convertedMaskPath = os.path.join(basePath, 'AP1_2D.nii.gz')
    sitk.WriteImage(mask, convertedMaskPath)
    maskFile = convertedMaskPath  # 更新掩膜路径

# 验证转换后的结果
print("Converted Image size:", image.GetSize())
print("Converted Image pixel type:", image.GetPixelIDTypeAsString())
print("Converted Mask size:", mask.GetSize())

# 初始化特征提取器
extractor = featureextractor.RadiomicsFeatureExtractor()

# 提取特征
try:
    featureVector = extractor.execute(imageFile, maskFile)
    for featureName, value in featureVector.items():
        print(f"{featureName}: {value}")
    featureDict = {featureName: value for featureName, value in featureVector.items()}

    df = pd.DataFrame([featureDict])
    outputFilePath = os.path.join(basePath, 'extracted_features.csv')
    df.to_csv(outputFilePath, index=False)
    print(f"Feature extraction completed and saved to {outputFilePath}")
except Exception as e:
    print(f"Error during feature extraction: {e}")
