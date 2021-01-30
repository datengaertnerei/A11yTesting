# Accessibility Testing goes Mainstream

## Links

### Standards und Verordnungen

- [Web Content Accessibility Guidelines 2.2](https://www.w3.org/TR/WCAG22/)
- [EU Richtlinie 2016/2102 über den barrierefreien Zugang zu den Websites und mobilen Anwendungen öffentlicher Stellen](https://eur-lex.europa.eu/legal-content/DE/TXT/HTML/?uri=CELEX:32016L2102&from=DE)
- [EU Richtlinie 2019/882 über die Barrierefreiheitsanforderungen für Produkte und Dienstleistungen](https://eur-lex.europa.eu/legal-content/DE/TXT/HTML/?uri=CELEX:32019L0882&from=DE)
- [Barrierefreie-Informationstechnik-Verordnung - BITV 2.0](https://www.gesetze-im-internet.de/bitv_2_0/BJNR184300011.html)

### Quellen

- [WAI: Evaluating Web Accessibility Overview](https://www.w3.org/WAI/test-evaluate/)
- [BITV Test](https://www.bitvtest.de/)
- [A11y Project](https://www.a11yproject.com/)
- [Using ARIA](https://www.w3.org/TR/using-aria/)
- [How to use NVDA and Firefox to test your web pages for accessibility](https://www.marcozehe.de/how-to-use-nvda-and-firefox-to-test-your-web-pages-for-accessibility/)
- [Netzwerk Leichte Sprache e.V.](https://www.leichte-sprache.org/)

### Tools

- [Google Lighthouse](https://developers.google.com/web/tools/lighthouse)
- [lighthouse-batch](https://github.com/mikestead/lighthouse-batch)
- [HTML Code Sniffer](http://squizlabs.github.io/HTML_CodeSniffer/)
- [Sa11y QA assistant](https://ryersondmp.github.io/sa11y/)
- [axe-core Browser Plugin](https://github.com/dequelabs/axe-core)
- [Pa11y](https://pa11y.org/)
- [NVDA Screen Reader](https://www.nvaccess.org/)
- [ESLint a11y plugin](https://www.npmjs.com/package/eslint-plugin-jsx-a11y)
- [Color Oracle](http://colororacle.org/)
- [Colour Contrast Analyser](https://developer.paciellogroup.com/resources/contrastanalyser/)

## Lighthouse Batch Spider

Im lighthouse Verzeichnis liegt die Konfiguration für eine automatisierte Lighthouse Auswertung einer Website. Dazu baut man sich einmalig ein passendes Docker Image:

```
docker build -t datengaertner/lhbatch .
```

oder lädt es vom Docker Hub:

```
docker pull datengaertner/lhbatch:latest
```

Danach ruft man den Batch so auf:

```
docker run --rm -e SITE="https://www.datengaertnerei.com/" datengaertner/lhbatch > summary.csv
```

Das Ergebnis ist eine CSV Tabelle mit den konsolidierten Lighthouse Ergebnissen.

## Pa11y Dashboard Docker setup

Im pally Verzeichnis stehen die Konfigurationsdateien für ein Pa11y Dashboard Docker Setup. Mit den Dateien im Verzeichnis kann man einfach die Container starten mit

```
docker-compose up -d
```

Der Dashboard Webserver ist über http Port 4000 erreichbar. Mi folgendem cURL Befehl bekommt man alle konfigurierten Tasks über den Webservice als JSON Export:

```
curl -X GET -o pally.json http://localhost:3000/tasks
```

Mit einem POST auf diesen Endpunkt kann man einzelne Tasks einspielen. 

```
curl -X POST -d @neu.json -H "Content-Type: application/json" http://localhost:3000/tasks
```

Dazu gibt es ein kleines Kommandozeilentool, mit dem der Export in einzelne Import Dateien aufgeteilt wird.
