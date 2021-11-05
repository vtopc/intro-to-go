# Comparing float implementations and JSON serialization

## input

`1234567890123456.87`

## output

```
float64: {"price":1234567890123456.8}
big.Float: {"price":"1.2345678901234568e+15"}
shopspring/decimal.Decimal: {"price":"1234567890123456.87"}
```
