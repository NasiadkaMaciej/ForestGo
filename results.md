Zarówno patrząc na wykresy jak i obserwując wizualizację symulacji można wysunąć pewne wnioski.

Stanowczym czynnikiem wpływającym na spalanie lasu jest jego gęstość.
Drzewa zapalaja się w tylko będąc obok płonącego drzewa (po skosie jest mniejsze prawdopodobieństwo). W związku z czym, las musi być odpowiednio gęsty, żeby spalił się w dużej części. Na wykresie burn_by_density można zaobserować znaczący skok przy gęstości ~0,4/0,5

Wilgotność wpływa odwrotnie proporcjonalnie do spalalności lasu - czym większa, tym mniejsza szansa na wielki pożar. Doskonale obrazuje to linia trendów na wykresie burn_by_humidity

Naturalnie, występują również wypadki specyficzne, jak np duża gęstość i wysoka wilgotność. Naturalnie, mimo dużej gęstości, drzewa raczej nie będą miały dużej skłonności do zapalenia się podczas intensywnych opadów.

UPDATE:
Po przekształceniu funkcji CreateForest (teraz wylicza na sztywno ilość drzew dla każdej gęstości) wykresy nie wyglądają "naturalnie"
Dane wygenerowane przez poprzednią funkcję dostępne są w pliku forest_fire_statsOldFun.csv. Zachęcam do wygenerowania i porównania wykresów.