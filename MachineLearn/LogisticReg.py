
import pandas as pd
import numpy as np
from sklearn.linear_model import LogisticRegression
from sklearn.model_selection import train_test_split, learning_curve, ShuffleSplit

def train_model_by_lr(file_path, regularization='l2'):
    # 读取数据
    data = pd.read_csv(file_path)
    X = data.iloc[:, :-1]
    y = data['label']

    # 划分训练集和测试集
    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.3, random_state=21)

    # 根据正则化参数创建模型
    if regularization == 'l1':
        model = LogisticRegression(penalty="l1", solver="liblinear", max_iter=1000)
    else:
        model = LogisticRegression(penalty="l2", solver="liblinear", max_iter=1000)

    model.fit(X_train, y_train)

    # 学习曲线
    cv = ShuffleSplit(n_splits=50, test_size=0.2, random_state=0)
    train_size_rate = np.linspace(0.1, 1, 10)
    train_sizes, train_scores, test_scores = learning_curve(model, X, y, cv=cv, train_sizes=train_size_rate)

    # 计算均值和标准差
    train_scores_mean = np.mean(train_scores, axis=1)
    train_scores_std = np.std(train_scores, axis=1)
    test_scores_mean = np.mean(test_scores, axis=1)
    test_scores_std = np.std(test_scores, axis=1)

    return {
        'model': model,
        'train_accuracy': model.score(X_train, y_train),
        'val_accuracy': model.score(X_test, y_test),
        'learning_curve': {
            'train_sizes': train_size_rate.tolist(),
            'train_scores_mean': train_scores_mean.tolist(),
            'train_scores_std': train_scores_std.tolist(),
            'test_scores_mean': test_scores_mean.tolist(),
            'test_scores_std': test_scores_std.tolist()
        }
    }




