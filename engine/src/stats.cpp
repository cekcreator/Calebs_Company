#include "stats.hpp"
#include <algorithm>
#include <cmath>
#include <numeric>
#include <stdexcept>

namespace engine {

Stats compute_stats(std::vector<double> data) {
    if (data.empty()) throw std::invalid_argument("data must not be empty");

    double mean = std::accumulate(data.begin(), data.end(), 0.0) / data.size();

    double variance = 0.0;
    for (double v : data) variance += (v - mean) * (v - mean);
    variance /= data.size();

    std::sort(data.begin(), data.end());
    double median = data.size() % 2 == 0
        ? (data[data.size() / 2 - 1] + data[data.size() / 2]) / 2.0
        : data[data.size() / 2];

    return { mean, variance, std::sqrt(variance), data.front(), data.back(), median };
}

} // namespace engine
