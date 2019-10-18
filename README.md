# af

Activation functions, intended for use in neural networks.

Provides the following activation functions, that take just one argument:

* Sigmoid (optimized, from the [swish](https://github.com/xyproto/swish) package).
* Swish (optimized, from the [swish](https://github.com/xyproto/swish) package).
* SoftPlus (optimized, from the [swish](https://github.com/xyproto/swish) package).
* Abs (`math.Abs`)
* Tanh (`math.Tanh`)
* Sin (`math.Sin`)
* Cos (`math.Cos`)
* Inv (`-x`)
* ReLU (`x > 0 ? x : 0`)

And also these functions, that take two arguments:

* PReLU (`x > 0 ? x : x * a`)

## General information

* License: MIT
* Version: 0.1.0
* Author: Alexander F. Rødseth &lt;xyproto@archlinux.org&gt;
