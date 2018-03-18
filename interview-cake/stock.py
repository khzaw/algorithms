def get_max_profit(stock_prices_yesterday):
    """
    Returns the best profit one could have made from
    1 purchase and 1 sale of stock yesterday
    """
    if len(stock_prices_yesterday) < 2:
        raise ValueError('Requires at least 2 prices')

    min_price = stock_prices_yesterday[0]
    max_profit = stock_prices_yesterday[1] - stock_prices_yesterday[0]

    for current_time in range(1, len(stock_prices_yesterday)):
        current_price = stock_prices_yesterday[current_time]

        profit = current_price - min_price
        max_profit = max(max_profit, profit)

        min_price = min(min_price, current_price)

    return max_profit


assert(get_max_profit([10, 7, 5, 8, 11, 9]) == 6)
assert(get_max_profit([10, 11, 22, 5, 2]) == 12)
