from sklearn.svm import SVC
import pandas as pd
from sklearn.model_selection import train_test_split

dataFile = r"E:\dachuang\Radiomics\Data\test_liang\extracted_features.xlsx"



data = pd.read_excel(dataFile)


X = data.iloc[: ,0:-1]
y = data['Label']
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size = 0.3,random_state=21)



svc = SVC(kernel='rbf',probability=True)


svc.fit(X_train,y_train)

print(svc.score(X_train,y_train))
print(svc.score(X_test,y_test))
print(svc.predict(X_test))
print(svc.predict_proba(X_test)[:,1])
