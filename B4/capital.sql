select stock_name, SUM(
    CASE 
        WHEN operation = "Sell" Then price 
        ELSE price * -1
    END
) as capital_gain_loss from Stocks 
group by stock_name