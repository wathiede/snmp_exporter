# Prometheus SNMP Exporter

Work in progress

## Usage

`config` contains what OIDs to scrape and how to process them. It initially
supports enough for the standard interface stats.

Run `snmp_exporter.py`, and then visit http://localhost:9116/?address=1.2.3.4
where 1.2.3.4 is the IP of the SNMP device to get metrics.

## Design

There are two components. An exporter that does the actual scraping,
and a generator that creates the configuration for use by the exporter.

This is to allow for customisation of what's done during the scrape as many special cases are expected.
The varying levels of SNMP MIB-parsing support across different langauges also
means that a single language may not be practical.t
