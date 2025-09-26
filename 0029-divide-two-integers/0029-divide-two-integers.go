func divide(dividend int, divisor int) int {
    quotient := 0
    sign := 1
    if dividend < 0 {
        dividend = -dividend
        sign *= -1
    }
    if divisor < 0 {
        divisor = -divisor
        sign *= -1
    }

    for dividend >= divisor {
        power := 1
        v := divisor
        for v + v  <= dividend {
            v += v
            power += power
        }
        quotient += power
        dividend -= v
    }

    quotient *= sign

    if quotient > 2147483647 {
        quotient = 2147483647
    }
    if quotient < -2147483648 {
        quotient = 2147483649
    }

    return quotient
}