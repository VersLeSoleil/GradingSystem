# app.py
from flask import Flask, request, jsonify, send_from_directory
import pandas as pd
import numpy as np
from sklearn.linear_model import LogisticRegression
from sklearn.model_selection import train_test_split, ShuffleSplit, learning_curve
from MachineLearn.LogisticReg import train_model_by_lr
from MachineLearn.RandomForest import train_model_by_rf
import joblib
import os
from PreProcessing.FeatureExtractPlus import extract_features
from flask_cors import CORS

app = Flask(__name__)
CORS(app)
# 全局变量存储模型和数据集
global_model = None
global_dataset = None
UPLOAD_FOLDER = 'uploads'
os.makedirs(UPLOAD_FOLDER, exist_ok=True)


@app.route("/")  # 确保有这个路由
def home():
    return "Hello, World!"  # 返回一些内容


@app.route('/index')
def serve_index():
    return send_from_directory('.', 'index.html')


@app.route('/upload', methods=['POST'])
def upload_file():
    print(f"收到上传请求，当前工作目录: {os.getcwd()}")
    if 'file' not in request.files:
        print('未找到文件！')
        return jsonify({'error': 'No file uploaded'}), 400

    file = request.files['file']
    print(f"接收到文件: {file.filename}")
    if file.filename == '':
        return jsonify({'error': 'No selected file'}), 400

    filepath = os.path.join(UPLOAD_FOLDER, file.filename)
    file.save(filepath)

    try:
        data = pd.read_csv(filepath)
        global global_dataset
        global_dataset = data
        return jsonify({
            'message': 'File uploaded successfully',
            'filename': file.filename,
            'columns': list(data.columns),
            'sample_size': len(data)
        })
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@app.route('/train', methods=['POST'])
def train_model_api():
    print("成功接收到训练请求")
    data = request.json
    model_type = data.get('model_type', 'logistic')
    regularization = data.get('regularization', 'l1')
    n_estimators = data.get('n_estimators', 100)  # 设置默认值
    max_depth = data.get('max_depth', 4)

    if global_dataset is None:
        return jsonify({'error': 'No dataset loaded'}), 400

    try:
        # 清洗 global_dataset，避免 NaN、inf 报错
        clean_df = global_dataset.copy()

        # 去除非数值型列（如 'Folder', 'Name' 等）
        clean_df = clean_df.select_dtypes(include=[np.number])

        # 替换 inf 为 NaN，再删除含空值的行
        clean_df.replace([np.inf, -np.inf], np.nan, inplace=True)
        clean_df.dropna(inplace=True)

        # 检查是否为空
        if clean_df.empty:
            return jsonify({'error': '训练数据为空，可能存在太多缺失值或非法值'}), 400

        # 保存临时文件用于训练
        temp_path = os.path.join(UPLOAD_FOLDER, "temp_dataset.csv")
        clean_df.to_csv(temp_path, index=False)

        # 训练模型
        if model_type == 'logistic':
            result = train_model_by_lr(temp_path, regularization=regularization)
        else:
            result = train_model_by_rf(temp_path, n_estimators=n_estimators, max_depth=max_depth)
        global global_model
        global_model = result['model']
        lc = result['learning_curve']
        # 兼容不同模型的字段
        learning_curve = {
            "labels": lc.get("estimators") or lc.get("train_sizes") or [],
            "train_scores": lc.get("train_scores") or lc.get("train_scores_mean") or [],
            "val_scores": lc.get("test_scores") or lc.get("test_scores_mean") or []
        }
        return jsonify({
            'status': 'success',
            'train_accuracy': result['train_accuracy'],
            'val_accuracy': result['val_accuracy'],
            'learning_curve': learning_curve,
            'confusion_matrix': result.get('confusion_matrix'),
            'feature_importance': result.get('feature_importance')

        })

    except Exception as e:
        print("❌ 训练异常：", str(e))
        return jsonify({'error': str(e)}), 500


@app.route('/predict_test', methods=['POST'])
def predict_test_api():
    if 'file' not in request.files:
        return jsonify({'error': 'No test file uploaded'}), 400

    test_file = request.files['file']
    if test_file.filename == '':
        return jsonify({'error': 'Empty filename'}), 400

    temp_test_path = os.path.join(UPLOAD_FOLDER, test_file.filename)
    test_file.save(temp_test_path)

    if global_model is None:
        return jsonify({'error': 'Model not trained'}), 400

    try:
        # ✅ 读取测试集
        test_data = pd.read_csv(temp_test_path, header=0)

        # ✅ 清洗数据：替换无穷大为 NaN，再删除缺失行
        test_data.replace([np.inf, -np.inf], np.nan, inplace=True)
        test_data.dropna(inplace=True)

        # ✅ 如果清洗后为空，说明全是问题数据
        if test_data.empty:
            return jsonify({'error': '清洗后测试集为空，可能包含大量缺失或异常值'}), 400

        # ✅ 模型预测
        preds = global_model.predict(test_data)
        probs = global_model.predict_proba(test_data)

        # ✅ 构建结果表格
        results = test_data.copy()
        results['predicted_label'] = preds
        for i in range(probs.shape[1]):
            results[f'prob_class_{i}'] = probs[:, i]

        return jsonify({
            'message': '预测成功',
            'data': results.to_dict(orient='records')
        })

    except ValueError as ve:
        return jsonify({'error': f'预测失败，输入数据中可能有非法值或格式不匹配：{str(ve)}'}), 400
    except Exception as e:
        return jsonify({'error': str(e)}), 500


@app.route('/extract_features', methods=['POST'])
def extract_features_api():
    try:
        data = request.get_json()
        basePath = data.get('basePath')
        outputPath = data.get('outputPath')
        if not basePath or not outputPath:
            return jsonify({'error': '参数不完整'}), 400
        msg = extract_features(basePath, outputPath)
        return jsonify({'message': msg})
    except Exception as e:
        print("❌ 特征提取异常：", str(e))
        return jsonify({'error': str(e)}), 500


if __name__ == '__main__':
    app.run(host="0.0.0.0", port=8080, debug=True)