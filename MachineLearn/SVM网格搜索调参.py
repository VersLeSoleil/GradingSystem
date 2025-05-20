from sklearn.svm import SVC
import pandas as pd
import numpy as np
from sklearn.model_selection import train_test_split
from sklearn.model_selection import GridSearchCV



dataFile = r"F:\Radiomics\Data\TestData\machineLearning\breast_cancer_m.csv"
data = pd.read_csv(dataFile)
X = data.iloc[: ,0:-1]
y = data['label']
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size = 0.3,random_state=21)



Cs = np.logspace(-1,3,10,base = 2)
gammas = np.logspace(-5,1,10)
param_grid = {'C': Cs, 'gamma': gammas, 'kernel': ('rbf','linear')}
GS = GridSearchCV(SVC()
                    , param_grid = param_grid
                    , cv = 5
                   )
GS.fit(X_train,y_train)
print(GS.best_params_)
C = GS.best_params_['C']
gamma = GS.best_params_['gamma']
kernel = GS.best_params_['kernel']

svc = SVC(kernel=kernel
          , C = C
          , gamma = gamma
          )
svc.fit(X_train,y_train)
print(svc.score(X_train,y_train))
print(svc.score(X_test,y_test))