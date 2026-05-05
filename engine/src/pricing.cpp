#include "pricing.hpp"
#include <stdexcept>

namespace engine {

PricingResult calculate_price(const PricingInput& in) {
    if (in.base_price < 0 || in.quantity <= 0)
        throw std::invalid_argument("base_price must be >= 0 and quantity > 0");
    if (in.discount_pct < 0 || in.discount_pct > 100)
        throw std::invalid_argument("discount_pct must be in [0, 100]");

    double subtotal        = in.base_price * in.quantity;
    double discount_amount = subtotal * (in.discount_pct / 100.0);
    double discounted      = subtotal - discount_amount;
    double tax_amount      = discounted * in.tax_rate;
    double total           = discounted + tax_amount;

    return { subtotal, discount_amount, tax_amount, total };
}

} // namespace engine
