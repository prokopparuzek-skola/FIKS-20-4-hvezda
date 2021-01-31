# Hvězda

## Nalezení všech frekvencí

Pro nalezení všech frekvencí jsem zvolil tento postup. Prvně si předpočítám prefixové součty pro všechny paprsky s tím 
že začínám u jádra za účelem zrychlení následných výpočtů. Poté pro každý paprsek spočítám všechny frekvence co se v něm 
nacházejí, mezi 1 a 3 částí, 2 a 3 ... a pro součty začínající u jádra přidám ještě to samé, avšak s přičtením jádra. 
A nakonec spočítám součty mezi dvěma paprsky přes jádro, vezmu všechny jejich kombinace a pro každou vytvořím dejme tomu 
pole začínající nejvzdálenější částí 1 paprsku pokračující přes jádro k nejzazšímu bodu druhého a spočítám to samé jako 
u samostatných paprsků, avšak s tím rozdílem, že vždy začínán v jednom paprsku a končím až v tom druhém. A samozřejmě 
přidám všechny hodnoty frekvencí, které načtu na vstupu.

Algoritmus projde všechny možné kombinace frekvencí, jak v paprscích tak mezi nimi a seřadí je asymptoticky myslím 
v čase $O(N^4)$. Prefixové součty spočítá v čase N^2^, N paprsků a N hodnot v každém, kombinace spočte v N^3^ pro každý 
bod spočte frekvence pro všechny následující. Avšak pro N paprsků. A pro mezi paprskové součty potřebuje N^4^, N^2^ 
kombinací paprsků a v každé kombinaci je dalších N^2^ frekvencí. Vše seřadím v čase $N*log(N)$. A po součtu a zanedbání 
nízkých mocnin zbude N^4^.

## Nalezení frekvence

Použil jsem v zásadě stejný algoritmus s malým kosmetickým zlepšením. Pokud v rámci součtů narazím na součet vyšší než 
požadovaná frekvence, můžu přejít na další část, protože součty už se budou jen zvyšovat. Ale asymptoticky to myslím 
pořád zůstává $O(N^4)$ a jde tedy jen o jak už jsem psal, asymptoticky bezvýznamné zrychlení.
