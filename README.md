określenie problemu
=============

Tworzymy sieć komputerową o strukturze drzewowej, dla każdej maszyny istnieje z góry określona maksymalna pojemność - ilość maszyn, z którymi może być połączona (zakładamy, że jest to liczba większa od 2). Chcemy zachować najmniejszą możliwą średnicę sieci, tj. chcemy, aby odległość między najbardziej oddalonymi wierzchołkami była jak najmniejsza. Dochodzi nowy komputer, który musimy dodać on-line. Co robimy?

propozycja rozwiązania
=============

IDEA ROZWIĄZANIA:
Tworzymy sieć w formie ukorzenionego drzewa (korzeń stanowi pierwsza maszyna). Nowy komputer dołączamy do takiego wierzchołka, który nie osiągnął maksymalnej pojemności i ma najniższą głębokość. Jeśli takich wierzchołków jest więcej, wybieramy ten najbardziej na lewo.

CO PRZEMAWIA ZA UŻYCIEM WŁAŚNIE TEGO ALGORYTMU DODANIA KOMPUTERA:
1. Nowo dodany komputer połączony będzie z co najwyżej jednym komputerem dodanym wcześniej. Dlatego nigdy nie powstanie nam cykl i sieć zachowa strukturę drzewa.
2. Każde drzewo ma liście. Pojemność wierzchołków jest większa od 2, zatem zawsze istnieje wierzchołek, do którego możemy coś dodać.
3. Zauważmy, że średnica powiększa się, gdy tworzymy nowy poziom w drzewie. Nasze postępowanie prowadzi jednak do tego, by nowy poziom był tworzony jak najrzadziej.

Oczywiście, globalnie rozwiązanie to nie jest optymalne, ale my musimy radzić sobie z problemem on-line!

JEDNAKŻE, POJAWIA SIĘ PEWIEN PROBLEM: co będzie, jeśli dostaniemy dwa requesty o dołączenie do sieci na raz? Możemy rozwiązać go w następujący sposób: root będzie znał strukturę drzewa.

PROPONOWANA DODATKOWA FUNKCJONALNOŚĆ: tworzenie obrazu naszej sieci na prośbę (ludzkich) użytkowników sieci. Korzeń, który zna strukturę sieci, tworzy wykres oraz go rozsyła.

szczegóły rozwiązania
=============

DODAWANIE NOWEGO KOMPUTERA: informację o tym, że przychodzi nowy komputer, otrzymuje dowolna maszyna w naszej sieci. Następnie każdy wysyła zapytanie o wynik do swojego rodzica, póki nie dojdzie do korzenia. Korzeń, używając wspomnianego wcześniej algorytmu, znajduje wierzchołek, do którego należy podpiąć nową maszynę. Rozsyła później informację do wszystkich swoich dzieci, które przekazują ją dalej. Wiadomość rozprzestrzenia się w całej sieci.

TWORZENIE WYKRESU: wierzchołek otrzymuje polecenie "zbuduj wykres". Wysyła wówczas do swojego rodzica tę wiadomość. Gdy wiadomość dochodzi do korzenia, tworzy on skrypt budujący wykres oraz rozsyła go do swoich dzieci z komunikatem mówiącym ZBUDOWAŁEM SKRYPT. Zostaje ona przekazana dalej.

szczegóły implementacji
=============

WYBRANY JĘZYK: go

POŁĄCZENIE MIĘDZY KOMUPUTERAMI: socket

POZOSTAŁE TECHNOLOGIE: Google Charts (https://developers.google.com/chart/)

wybrane protokoły
=============

PROTOKÓŁ WARSTWY TRANSPORTOWEJ: tcp. Nie zależy nam na szybkości, natomiast przy tworzeniu sieci bardzo ważne są dla nas bezpieczeństwo oraz dokładność.

PROTOKÓŁ 1 - PRZESYŁANIE INFORMACJI MIĘDZY KOMPUTERAMI (podczas tworzenia sieci, przekazywania informacji o klientach, "dogadywania się" w sprawie przesłania skryptu wykresu): wystarczy bardzo prosty protokół tekstowy. Jedyne informacje, jakie potrzebujemy, to: typ wiadomości, jej treść oraz ewentualna informacja o błędzie. Zauważmy, że jest to wiadomość, która jest bardzo niewielka, nie potrzebujemy zatem w żaden sposób "porcjować" danych. Będzie ona przekazywana w taki sposób, że wysyłamy tę samą wiadomość póki nie otrzymamy informacji zwrotnej - potwierdzenia jej otrzymania.

Protokołowi temu zdecydowaliśmy nadać nazwę Simple Information Protocol - w skrócie SIP.

PROTOKÓŁ 2 - PRZESYŁANIE WYKRESU: plik z wykresem może być duży, musimy zatem przesyłać go w mniejszych pakietach. Możemy użyć protokołu tekstowego i przekazać po prostu kod wykresu w JavaScripcie. Zauważmy, że komputery mogą ustalić między sobą fakt, że zaraz nastąpi tranfser danych za pomocą naszego protokołu. Następnie możemy podzielić przesyłany skrypt na odpowiednio małe fragmenty i przesyłać je do momentu otrzymania informacji zwrotnej na ich temat.

Protokół ten nazwaliśmy Script Transfet Protocol - STP.

Dlaczego zamiast używania SIP oraz STP nie użyjemy jednego protokołu, o formacie jak SIP, skoro równie dobrze można wysyłać podzielony plik za pomocą SIP? Zwróćmy uwagę na to, ile niepotrzebnych informacji zostałoby wówczas przekazane i jak bardzo spowolniłoby to transfer danych!

dowcip
=============

Czego potrzeba do ukończenia projektu z sieci?

Szczypki kreatywności.