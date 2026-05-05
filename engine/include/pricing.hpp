#pragma once
#include <vector>

namespace engine {

struct PricingInput {
    double base_price;
    double quantity;
    double discount_pct;   // 0.0 – 100.0
    double tax_rate;       // 0.0 – 1.0
};

struct PricingResult {
    double subtotal;
    double discount_amount;
    double tax_amount;
    double total;
};

PricingResult calculate_price(const PricingInput& input);

} // namespace engine
