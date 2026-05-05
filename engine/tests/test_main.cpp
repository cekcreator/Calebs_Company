#include <cassert>
#include <cmath>
#include <iostream>
#include "pricing.hpp"
#include "stats.hpp"

void test_pricing() {
    engine::PricingInput in{ 100.0, 2.0, 10.0, 0.1 };
    auto r = engine::calculate_price(in);
    assert(std::abs(r.subtotal - 200.0) < 1e-9);
    assert(std::abs(r.discount_amount - 20.0) < 1e-9);
    assert(std::abs(r.total - 198.0) < 1e-9);
    std::cout << "PASS: pricing\n";
}

void test_stats() {
    auto s = engine::compute_stats({ 1.0, 2.0, 3.0, 4.0, 5.0 });
    assert(std::abs(s.mean - 3.0) < 1e-9);
    assert(std::abs(s.median - 3.0) < 1e-9);
    assert(s.min == 1.0 && s.max == 5.0);
    std::cout << "PASS: stats\n";
}

int main() {
    test_pricing();
    test_stats();
    std::cout << "All tests passed.\n";
    return 0;
}
