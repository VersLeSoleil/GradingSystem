import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.preprocessing import StandardScaler
from sklearn.svm import SVC
from sklearn.metrics import classification_report, confusion_matrix


data = pd.read_excel("extracted_features.xlsx")

# 提取特征和标签
X = data.iloc[:, 1:].values  # 特征
y = data.iloc[:, 0].values   # 标签（第一列为Label）


X = pd.DataFrame(X).fillna(0).values  # 将缺失值填充为0

# 数据标准化
scaler = StandardScaler()
X_scaled = scaler.fit_transform(X)

# 分割数据集（训练集 80%，测试集 20%）
X_train, X_test, y_train, y_test = train_test_split(X_scaled, y, test_size=0.2, random_state=42, stratify=y)

# 创建并训练 SVM 模型
svm_model = SVC(kernel='rbf', random_state=42)
svm_model.fit(X_train, y_train)

# 模型预测
y_pred = svm_model.predict(X_test)

# 评估模型
print("分类报告:")
print(classification_report(y_test, y_pred))

print("混淆矩阵:")
print(confusion_matrix(y_test, y_pred))
