import statsmodels.api as sm
import statsmodels.formula.api as smf
import matplotlib.pyplot as plt
import numpy as np
import pandas as pd
from patsy import dmatrices
from sklearn.linear_model import LinearRegression

# column_names = ['CRIM', 'ZN', 'INDUS', 'CHAS', 'NOX', 'RM', 'AGE', 'DIS', 'RAD',
#                 'TAX', 'PTRATIO', 'B', 'LSTAT']
filename = '/home/run/workspace/go/src/FinancialAnalysis/BRK.B.csv'

df = pd.read_csv(filename, parse_dates=[0], index_col='timestamp')
df = df.sort_index();
print(df.head())

prices = df['close'].tolist()
dates = df.index.values.astype(float).tolist()
 
#Convert to 1d Vector
dates = np.reshape(dates, (len(dates), 1))
prices = np.reshape(prices, (len(prices), 1))

print(dates)
#Define Linear Regressor Object
regressor = LinearRegression()
regressor.fit(dates, prices)
 
#Visualize Results
plt.scatter(dates, prices, color='yellow', label= 'Actual Price') #plotting the initial datapoints
plt.plot(dates, regressor.predict(dates), color='red', linewidth=3, label = 'Predicted Price') #plotting the line made by linear regression
plt.title('Linear Regression | Time vs. Price')
plt.legend()
plt.xlabel('Date Integer')
plt.show()
 
#Predict Price on Given Date
date = 1536816743713000
predicted_price =regressor.predict(date)
print(predicted_price)
print(predicted_price[0][0],regressor.coef_[0][0] ,regressor.intercept_[0])