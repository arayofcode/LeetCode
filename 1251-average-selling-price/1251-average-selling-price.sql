SELECT Prices.product_id, 
    CASE WHEN SUM(units) is NULL 
        THEN 0
        ELSE ROUND(SUM(units * price)/ SUM(units)::numeric, 2) 
    END average_price
FROM Prices 
LEFT JOIN UnitsSold 
    ON Prices.product_id = UnitsSold.product_id
    AND purchase_date >= start_date 
    AND purchase_date <= end_date
GROUP BY Prices.product_id