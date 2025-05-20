# radiomics_extract.py
from radiomics import featureextractor
import os
import SimpleITK as sitk
import pandas as pd

def extract_features(basePath, outputPath):
    extractor = featureextractor.RadiomicsFeatureExtractor()
    all_features = []

    for subfolder in os.listdir(basePath):
        subfolderPath = os.path.join(basePath, subfolder)
        if os.path.isdir(subfolderPath):
            imageFile = os.path.join(subfolderPath, 'AP1.jpg')
            maskFile = os.path.join(subfolderPath, 'AP1.nii.gz')
            print(f"处理: {subfolder}")
            if not os.path.exists(imageFile):
                print(f"缺失影像: {imageFile}")
                continue
            if not os.path.exists(maskFile):
                print(f"缺失掩膜: {maskFile}")
                continue

            
    # 读取jpg并保存为nii.gz
            try:
                image = sitk.ReadImage(imageFile)
                nii_image_path = os.path.join(subfolderPath, 'AP1_converted.nii.gz')
                sitk.WriteImage(image, nii_image_path)
                imageFile = nii_image_path  # 后续都用nii.gz
            except Exception as e:
                print(f"JPG转NII失败: {imageFile}, 错误: {e}")
                continue

            try:
                mask = sitk.ReadImage(maskFile)
            except Exception as e:
                print(f"掩膜读取失败: {maskFile}, 错误: {e}")
                continue

            if image.GetNumberOfComponentsPerPixel() > 1:
                image = sitk.VectorIndexSelectionCast(image, 0)
                imageFile = os.path.join(subfolderPath, 'AP1_gray.nii.gz')
                sitk.WriteImage(image, imageFile)

            if len(mask.GetSize()) == 3 and mask.GetSize()[-1] == 1:
                mask = sitk.Extract(mask, mask.GetSize()[:-1] + (0,))
                maskFile = os.path.join(subfolderPath, 'AP1_2D.nii.gz')
                sitk.WriteImage(mask, maskFile)

            try:
                featureVector = extractor.execute(imageFile, maskFile)
                featureDict = {k: v for k, v in featureVector.items()}
                featureDict['Folder'] = subfolder
                all_features.append(featureDict)
            except Exception as e:
                print(f"Error: {e}")

    if all_features:
        df = pd.DataFrame(all_features)
        df.to_csv(outputPath, index=False)
        return f"特征提取成功，共提取 {len(df)} 条记录，保存到 {outputPath}"
    else:
        return "未提取到任何特征，请检查数据是否完整。"
