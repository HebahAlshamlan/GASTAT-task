# GASTAT-task
Case interview for Data Scientist


# Dataset
Tamimi markets products dataset

### Metadata

| # of records | 16K |
| ------- | --- | 
| Collecting Time | May2020 |
|  Max Price | 699.95 | 
| Min Price | 0.25 |
| Avg Price | 22.8 |
|# of Features | 11 |
#### Features

- Created time (Date)
- brand (Nominal)
- Product Name (Nominal)
- Product Variant (Nominal)
- Category (Nominal)
- General Category (Nominal)
- Class (Nominal)
- Super Class (Nominal)
- Discount (Numeric)
- VAT (Numeric)
- Price (Numeric)

# Consumer Price Index (CPI)
is a measure that examines the weighted average of prices of a basket of consumer goods and services, such as transportation, food, and medical care. It is calculated by taking price changes for each item in the predetermined basket of goods and averaging them. Changes in the CPI are used to assess price changes associated with the cost of living. The CPI is one of the most frequently used statistics for identifying periods of inflation or deflation.

#### The formula: sum(this month / last month)*100

# Monthly change(Laspeyres formula)
Basically it calculate the change rate between two period of time

#### The formula : (Production(this month/ last month)^item weight )^ (1/ Group wight)


# Pipeline 
for simplicity I try to collect the data only from Timimi market, because collecting data from multiple resorces and trying to do either Web Scraping or Use their API's it's a time-consuming.
- [X] Read and understand the concept
- [X] collect the data for April and March
- [X] Preprocssing 
  - [X] Cleaning
    - [X] Noise data
    - [X] Misspelling 
    - [X] Duplicates
    - [X] Outliers
   - [X] Missing Value
      - [X] Think of the best approch &Read the Literature
      - [X] Search for a reliable source
      - [X] Collecteing manually
 - [X] implementation
    - [X] CPI formula
    - [X] Monthly change
  
# Outcomes (findings)
- Customer price index of April 2020
- The percentage change in the index section Housing and its descendants
- The percentage change in the index section food and beverage

# Technologies and languages used
- Gathering data :
  - Go
  - Python
- Preprocessing
  - Excel
  - Weka
  - Python
- Implementation
  - Python
- ML Models and plots
  - Weka
