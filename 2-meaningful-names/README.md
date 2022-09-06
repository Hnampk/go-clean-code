# Chapter 2: Meaningful names

A summarize of chapter 2


1. intention-revealing names
2. avoid disinformation
    - poor variables names like `hp`, `aix`, `sco` because they are the name of Unix platforms or variants
    - it's probably better not to include the type into the name: `accounts/accountGroup/bunchOfAccounts` is better than `accountList` :)
    - beware of using names which vary in small ways: `XYZControllerForEfficientHandlingOfStrings` and `XYZControllerForEfficientStorageOfStrings`
    - careful on using lowercase of `L`, and uppercase of `o`: `l` looks likes `1`, and `O` looks like `0`
3. Make meaningful distinctions
    - number-series naming `(a1, a2,..., aN)` is opposite of intentional naming
    - noise words are meaningless distinction 
        + imagine you have `Product`, `ProductInfo`, `ProductData` classes, then what is the different?
        + one class named `Customer` and another named `CustomerObject` :)
        + how about a project has these function: `getActiveAccount`, `getActiveAccounts` and `getActiveAccountInfo`? Which should programers call?
    - sample of indistinguishable variable names
        + `moneyAmount` vs `money`
        + `customerInfo` vs `customer`
        + `accountData` vs `account`
        + `theMessage` vs `message`
4. use Pronounceable name
    - genymdhms (generate year month date hour minutes seconds), how to pronouce it? => generateTimestamp :)