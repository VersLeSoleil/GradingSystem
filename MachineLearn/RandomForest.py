import pandas as pd
import numpy as np
from sklearn.ensemble import RandomForestClassifier
from sklearn.model_selection import train_test_split
from sklearn.metrics import confusion_matrix

def train_model_by_rf(file_path, n_estimators=100, max_depth=None, max_estimators=50, test_size=0.3, random_state=21):
    data = pd.read_csv(file_path)
    X = data.iloc[:, :-1]
    y = data['label']

    X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=test_size, random_state=random_state)

    # ✅ 使用传入的 n_estimators 和 max_depth 训练最终模型
    clf_rf = RandomForestClassifier(n_estimators=n_estimators, max_depth=max_depth, random_state=random_state)
    clf_rf.fit(X_train, y_train)

    # 📈 生成学习曲线（用于可视化 n_estimators 从 1 到 max_estimators 的训练情况）
    scoreTrainList, scoreTestList = [], []
    for i in range(1, max_estimators + 1):
        temp_model = RandomForestClassifier(n_estimators=i, max_depth=max_depth, random_state=random_state)
        temp_model.fit(X_train, y_train)
        scoreTrainList.append(temp_model.score(X_train, y_train))
        scoreTestList.append(temp_model.score(X_test, y_test))

    best_score = max(scoreTestList)
    best_n = scoreTestList.index(best_score) + 1
    print(f"Best Test Accuracy: {best_score:.4f} at n_estimators = {best_n}")

    # 混淆矩阵
    y_pred = clf_rf.predict(X_test)
    cm = confusion_matrix(y_test, y_pred)

    # 特征重要性
    feature_importance = [
        {'name': name, 'importance': float(imp)}
        for name, imp in zip(X.columns, clf_rf.feature_importances_)
    ]
    feature_importance.sort(key=lambda x: x['importance'], reverse=True)

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
        },
        'confusion_matrix': {
            'true_negative': int(cm[0][0]),
            'false_positive': int(cm[0][1]),
            'false_negative': int(cm[1][0]),
            'true_positive': int(cm[1][1])
        },
        'feature_importance': feature_importance
    }