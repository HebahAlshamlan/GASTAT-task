# GASTAT-task
Case interview for Data Scientist


Not completed yet[Soon]

# Dataset
Tamimi markets products

### Metadata

| # of records | 16K |
| ------- | --- | 
| Collecting Time | May2020 |
|  Max Price | 699.95 | 
| Min Price | 0.25 |
| Avg Price | 22.8 |
|# of fetchers | 11 |
#### fetchers

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
Basically it's calculate the change rate between two period of time

#### The formula : (Production(this month/ last month)^item weight )^ (1/ Group wight)


# Pipeline 
for simplicity I try to collect the data only from Timimi market, becouse collecting data from multiple resorces and try to do ether Web Scraping or Use their API's time-consuming. 
- [X] coullect the data for April and March
- [X] Preprocssing 
  - [X] Cleaning
    - [X] Noise data
    - [X] Misspell
    - [X] Duplicates
    - [X] Outliers
   - [X] Missing Value
      - [X] Think of the best approch &Read the litrecher
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
  - python
- implementation
  - python
- ML Models and plots
  - Weka
