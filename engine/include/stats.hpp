#pragma once
#include <vector>

namespace engine {

struct Stats {
    double mean;
    double variance;
    double stddev;
    double min;
    double max;
    double median;
};

Stats compute_stats(std::vector<double> data);

} // namespace engine
