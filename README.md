My house's power-company, Georgia Power, a division of 
[Southern Company](https://en.wikipedia.org/w/index.php?title=Southern_Company&oldid=1197773773),
does a terrible job at informing me if I should be using a different rate plan.

They claim that my power meter doesn't track the usage well enough. But, from
the web site, I can download what appears to be hourly usage. I think they're
splitting hairs and intentionally crippling the power meter so that it doesn't
track *exactly* the same way, and therefore can claim that they can't give me
an estimate because they can't give me an exact value.


Why Is This Project Here?
=========================

No mortal can compare different rate plans for Southern Company.

The residential rate plan has a set of rules for summer and for non-summer.
The summer rules have tiers of cost, based on how much you have used since
the beginning of the billing period. First 600 kWh are $X each, and the next
350 are $Y each. There's a third tier too. And there is a fuel fee and a half
dozen other fees.

The time-of-use PEV rate has a set of rules for summer and for non-summer,
and they are the same except the summer has an additional set of rules for
the busiest time of day. There is no usage tier system. The cost for 
super-off-peak is $X each and it's at these hours. The cost for on-peak is $Y
and it's these hours but this condition doesn't exist during the non-summer or
during weekends. Anything not in the first, or (season-dependent) the second,
falls under off-peak, which is $Z each. And this has a different per-ruleset
fuel fee, but the same other fees.

Southern Company pretends that this is something a regular customer should
be able to make decisions about.


How Do I Use It?
================

Log in to Georgia Power's web site.

Visit "Billing and Payments" / "Usage" and download a month of *Hourly* usage
data.

https://customerservice2.southerncompany.com/Billing/MyPowerUsage

and that will save a "xlsx" file on disk. Load that directly into this program

	`go run ./cli.go ~/Downloads/GPC_Usage_2023-12-12-2024-01-11.xlsx`

and you should see something like

```
from file GPC_Usage_2023-12-12-2024-01-11.xlsx
             GP R-28: $64.60  Residential
       GP TOU-PEV-12: $78.81  Time-of-Use Plug-in Electric Vehicle
       GP TOU-REO-16: $74.63  Time-of-Use Residential Energy Only
```


Help
====

This is a work in progress. Bug reports welcome at github.
