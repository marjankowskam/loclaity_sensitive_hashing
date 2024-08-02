#pip install numpy pandas seaborn matplotlib

import numpy as np
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt
import csv

def load_arrays_from_csv(filename):
    array1 = []
    array2 = []
    array3 = []

    with open(filename, 'r') as file:
        reader = csv.reader(file)
        for row in reader:
            array1.append(float(row[0]))
            array2.append(float(row[1]))
            array3.append(float(row[2]))

    return array1, array2, array3

maxima, average, minima = load_arrays_from_csv("data/similarity_array_lineplot.csv")
x_labels = list(range(1, len(maxima) + 1))

plt.plot(x_labels, maxima, label='Max', marker='o', linestyle="--", color="blue")
plt.plot(x_labels, average, label='Mean', marker='s', color="blue")
plt.plot(x_labels, minima, label='Min', marker='^', linestyle="-.", color="blue")

plt.legend()
plt.title("Similarity after perturbing one of the flags/arguments")
plt.xlabel("Total number of flags and args in the sequence")
plt.ylabel("LSH similarity over 100 runs")

plt.grid(True, axis='y')
plt.xticks(x_labels)

plt.show()


plt.savefig('data/similarity_heatmap.png')
