#pip install numpy pandas seaborn matplotlib

import numpy as np
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt


grayscale = True
cmap = 'gray' if grayscale else 'viridis'

# Read the CSV file into a DataFrame
df = pd.read_csv('similarity_array.csv')

# Convert the DataFrame to a NumPy array
data = df.values

print('Hello')

# Plot the heatmap
plt.figure(figsize=(10, 8))

show_numbers = len(data) > 1000
sns_plot = sns.heatmap(data, annot=show_numbers, cmap=cmap)
plt.title('Heatmap of CSV Data')
plt.show()
#plt.savefig()

fig = sns_plot.get_figure()
fig.savefig('similarity_heatmap.png')
