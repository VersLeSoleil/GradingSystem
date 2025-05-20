import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from sklearn.ensemble import RandomForestClassifier
from sklearn.model_selection import train_test_split

def train_model_by_rf(file_path, max_estimators=50, test_size=0.3, random_state=21):
    """
    用随机森林模型训练数据并绘制学习曲线。

    :param file_path: 含有特征和标签的CSV路径，标签列为 'Label'
    :param max_estimators: 最大的树个数（用于绘制学习曲线）
    :param test_size: 测试集划分比例
    :param random_state: 随机种子
    :return: 包含模型对象、训练准确率、测试准确率、学习曲线信息的字典
    """
    # 读取数据
    data = pd.read_csv(file_path)
    X = data.iloc[:, :-1]
    y = data['label']

    # 划分训练集与测试集
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=test_size, random_state=random_state)

    # 训练最终模型（用最佳参数或者默认参数）
    clf_rf = RandomForestClassifier(n_estimators=max_estimators, random_state=random_state)
    clf_rf.fit(X_train, y_train)

    # 记录每个估计器数量下的训练/测试准确率
    scoreTrainList, scoreTestList = [], []
    for i in range(1, max_estimators + 1):
        temp_model = RandomForestClassifier(n_estimators=i, random_state=random_state)
        temp_model.fit(X_train, y_train)
        scoreTrainList.append(temp_model.score(X_train, y_train))
        scoreTestList.append(temp_model.score(X_test, y_test))

    # 输出最优估计器数量
    best_score = max(scoreTestList)
    best_n = scoreTestList.index(best_score) + 1
    print(f"Best Test Accuracy: {best_score:.4f} at n_estimators = {best_n}")

    return {
        'model': clf_rf,
        'train_accuracy': clf_rf.score(X_train, y_train),
        'val_accuracy': clf_rf.score(X_test, y_test),
        'learning_curve': {
            'estimators': list(range(1, max_estimators + 1)),
            'train_scores': scoreTrainList,
            'test_scores': scoreTestList,
            'best_n_estimators': best_n,
            'best_test_score': best_score
        }
    }


