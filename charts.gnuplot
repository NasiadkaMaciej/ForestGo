set datafile separator ","

# Common settings for all charts
set terminal pngcairo size 1000,700 font "Arial,12" enhanced
set grid xtics ytics mxtics mytics lt 1 lc rgb "#dddddd", lt 0 lc rgb "#eeeeee"
set key outside top right box
set style line 1 lc rgb "#00cc00" pt 7 ps 0.8 # green
set style line 2 lc rgb "#006600" lw 3        # dark green

# Burn Percentage vs. Forest Density
set output "simulation_results/burn_by_density.png"
set title "Burn Percentage vs. Forest Density" font "Arial,16"
set xlabel "Forest Density" font "Arial,14"
set ylabel "Burn Percentage (%)" font "Arial,14"
set xrange [0:1]
set yrange [0:*]

plot "simulation_results/forest_fire_stats.csv" using 1:5 with points ls 1 title "Data Points", \
     "" using 1:5 smooth acsplines ls 2 lw 3 title "Adaptive Spline Trend"
unset output

# Simulation Steps vs. Forest Density
set output "simulation_results/steps_by_density.png"
set title "Simulation Steps vs. Forest Density" font "Arial,16"
set xlabel "Forest Density" font "Arial,14"
set ylabel "Simulation Steps" font "Arial,14"
set xrange [0:1]
set yrange [0:*]

plot "simulation_results/forest_fire_stats.csv" using 1:6 with points ls 1 title "Data Points", \
     "" using 1:6 smooth acsplines title "Adaptive Spline Trend" ls 2 lw 3
unset output

# Burn Percentage vs. Humidity
set output "simulation_results/burn_by_humidity.png"
set title "Burn Percentage vs. Humidity" font "Arial,16"
set xlabel "Humidity" font "Arial,14"
set ylabel "Burn Percentage (%)" font "Arial,14"
set xrange [0:1]
set yrange [0:*]

plot "simulation_results/forest_fire_stats.csv" using 2:5 with points ls 1 title "Data Points", \
     "" using 2:5 smooth acsplines title "Adaptive Spline Trend" ls 2 lw 3
unset output

# Simulation Steps vs. Humidity
set output "simulation_results/steps_by_humidity.png"
set title "Simulation Steps vs. Humidity" font "Arial,16"
set xlabel "Humidity" font "Arial,14"
set ylabel "Simulation Steps" font "Arial,14"
set xrange [0:1]
set yrange [0:*]

plot "simulation_results/forest_fire_stats.csv" using 2:6 with points ls 1 title "Data Points", \
     "" using 2:6 smooth acsplines title "Adaptive Spline Trend" ls 2 lw 3
unset output