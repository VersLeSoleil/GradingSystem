import os
import pandas as pd
from sklearn.preprocessing import StandardScaler
from sklearn.svm import SVC
from sklearn.model_selection import train_test_split
from sklearn.metrics import accuracy_score
from sklearn.decomposition import PCA


basePath = r'F:\Radiomics\Data\Test'


outputFilePath = os.path.join(basePath, 'extracted_features.csv')
df = pd.read_csv(outputFilePath)


print("列名列表：", df.columns)


if 'Folder' in df.columns:
    df = df.drop(columns=['Folder'])
else:
    print("没有名为 'Folder' 的列。")


print(df.dtypes)


X = df.select_dtypes(include=['float64', 'int64'])


if X.empty:
    print("没有找到数值型数据列。")
else:

    X = X.fillna(X.mean())  # 用列的均值填充缺失值

    # 标准化特征
    scaler = StandardScaler()
    X_scaled = scaler.fit_transform(X)

    # 执行PCA降维
    pca = PCA(n_components=49)
    X_pca = pca.fit_transform(X_scaled)


    # 如果没有标签，你可以使用其他方法来创建标签，比如基于文件夹的分类
    y = df['label']  # 你需要根据实际情况设置标签

    # 划分数据集为训练集和测试集
    X_train, X_test, y_train, y_test = train_test_split(X_scaled, y, test_size=0.2, random_state=42)


    svm = SVC(kernel='linear')
    svm.fit(X_train, y_train)


    y_pred = svm.predict(X_test)
    print(svm.score(X_train, y_train))
    print(svm.score(X_test, y_test))
    accuracy = accuracy_score(y_test, y_pred)
    print(f"SVM分类准确率: {accuracy:.4f}")

    # 如果需要查看PCA的方差解释比例，可以使用以下代码：
    print("PCA方差解释比例：", pca.explained_variance_ratio_)
