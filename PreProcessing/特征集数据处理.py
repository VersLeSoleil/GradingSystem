from sklearn.preprocessing import MinMaxScaler
from sklearn.model_selection import train_test_split
import pandas as pd



dataFile = r"F:\Radiomics\Data\TestData\machineLearning\breast_cancer_m.csv"
data = pd.read_csv(dataFile)
data.head()
X = data.iloc[: ,0:-1]
y = data['label']
print(X.head())

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size = 0.3)
scaler_mm = MinMaxScaler()
scaler_mm_train = scaler_mm.fit(X_train)

X_train_mm = scaler_mm_train.transform(X_train)

X_train_mm = pd.DataFrame(X_train_mm,columns = X_train.columns)
X_test_mm = scaler_mm_train.transform(X_test)
X_test_mm = pd.DataFrame(X_test_mm,columns = X_test.columns)
X_test_mm2 = scaler_mm.fit_transform(X_test)
X_test_mm2 = pd.DataFrame(X_test_mm2,columns = X_test.columns)

scaler_mm2 = MinMaxScaler(feature_range = (-1,1))
scaler_mm2_train = scaler_mm2.partial_fit(X_train)
X_train_mm2 = scaler_mm2_train.transform(X_train)
X_train_mm2 = pd.DataFrame(X_train_mm2,columns = X_train.columns)

from sklearn.preprocessing import StandardScaler

scaler_ss = StandardScaler()
scaler_ss_train = scaler_ss.fit(X_train)
X_train_ss = scaler_ss_train.transform(X_train)
X_train_ss = pd.DataFrame(X_train_ss, columns = X_train.columns)

scaler_ss = StandardScaler()
X_train_ss = scaler_ss.fit_transform(X_train)
X_test_ss = scaler_ss.transform(X_test)
X_test_ss = scaler_ss.fit_transform(X_test)

