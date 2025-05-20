import pydicom
import os


dcmFolder = '/Users/lirenyuan/Desktop/maitian/course/data/mri/pat0001'
dcmFile = os.path.join(dcmFolder ,'0003_0001_568370669.dcm')
ds = pydicom.dcmread(dcmFile)
fileList = os.listdir(dcmFolder)
folder_new = '/Users/lirenyuan/Desktop/RadiomicsWorld/course/data/mri/pat0001_de'

if not os.path.exists(folder_new):
    os.makedirs(folder_new)
for file in fileList:
    filePath = os.path.join(dcmFolder, file)
    ds = pydicom.dcmread(filePath)
    ds.PatientName = 'XXX'
    ds.PatientID = '0000'
    ds.save_as(os.path.join(folder_new,file))
