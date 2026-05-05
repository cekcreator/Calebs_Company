#include <iostream>
#include <sstream>
#include "pricing.hpp"
#include "stats.hpp"

int main() {
    // Demo: pricing calculation
    engine::PricingInput input{ .base_price = 100.0, .quantity = 3.0, .discount_pct = 10.0, .tax_rate = 0.08 };
    auto result = engine::calculate_price(input);
    std::cout << "Pricing demo:\n"
              << "  subtotal:        " << result.subtotal        << "\n"
              << "  discount_amount: " << result.discount_amount << "\n"
              << "  tax_amount:      " << result.tax_amount      << "\n"
              << "  total:           " << result.total           << "\n\n";

    // Demo: stats calculation
    std::vector<double> data = { 4.0, 7.0, 13.0, 2.0, 1.0 };
    auto stats = engine::compute_stats(data);
    std::cout << "Stats demo:\n"
              << "  mean:    " << stats.mean    << "\n"
              << "  stddev:  " << stats.stddev  << "\n"
              << "  median:  " << stats.median  << "\n"
              << "  min/max: " << stats.min << " / " << stats.max << "\n";

    return 0;
}
