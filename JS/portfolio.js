import { Money } from './money.js'

export class Portfolio {
    constructor() {
        this.moneys = [];
    }

    add(...money) {
        this.moneys = this.moneys.concat(money);
    }

    evaluate(currency) {
        let total = this.moneys.reduce( (sum, money) => {
            return sum + this.convert(money, currency);
        }, 0);
        return new Money(total, currency);
    }
    convert(money, currency) {
        let eurToUsd = 1.2;
        if(money.currency === currency)
            return money.amount;
        return money.amount * eurToUsd;
    }
}