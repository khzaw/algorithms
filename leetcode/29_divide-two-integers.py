class Solution:
    INT_MIN: int = -(2**31)
    INT_MAX: int = 2**31 - 1

    def divide(self, dividend: int, divisor: int) -> int:
        sign = -1 if (dividend < 0) ^ (divisor < 0) else 1

        curr_dividend, q = abs(dividend), 0
        curr_divisor = abs(divisor)

        while curr_dividend >= curr_divisor:
            temp_d, temp_q = curr_divisor, 1

            while curr_dividend >= temp_d + temp_d:
                temp_d += temp_d
                temp_q += temp_q

            curr_dividend -= temp_d
            q += temp_q

        result = q * sign

        return max(min(result, self.INT_MAX), self.INT_MIN)
