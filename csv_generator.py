import csv
import random

# Define the CSV file path
csv_file_path = 'sample_data_analysis_8.csv'

name_arr = ["David", "Jack", "Sarah", "Jane", "Tom", "Jerry", "John", "Doe", "Smith", "Brown", "Kim", "Park", "Choi", "Jung", "Han", "Kang"]
category_arr = ["Luxury", "Bills", "Grocery", "Entertainment", "Travel"]

# Generate CSV data with 1000 lines
header = ['ID', 'Name', 'Credit', 'Category']
rows = [[i, random.choice(name_arr), random.randint(1, 10000), random.choice(category_arr)] for i in range(1, 100001)]



# Create and write data to the CSV file
with open(csv_file_path, mode='w', newline='') as file:
    writer = csv.writer(file)
    writer.writerow(header)
    writer.writerows(rows)

print(f"CSV file created at {csv_file_path}")
