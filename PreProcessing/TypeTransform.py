import SimpleITK as sitk



folderPath = '/Users/lirenyuan/Desktop/maitian/course/data/mri/pat0001_de'
reader = sitk.ImageSeriesReader()
dicom_names = reader.GetGDCMSeriesFileNames(folderPath)
reader.SetFileNames(dicom_names)
image = reader.Execute()
sitk.WriteImage(image,folderPath + '.nii') #这里亦可转换为压缩文件 .nii.gz
size = image.GetSize() # order: x, y, z
origin = image.GetOrigin() # order: x, y, z
spacing = image.GetSpacing() # order:x, y, z
direction = image.GetDirection() # order: x, y, z
imageArr = sitk.GetArrayFromImage(image) # order:z, y, x