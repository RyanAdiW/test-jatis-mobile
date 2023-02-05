SELECT extract(YEAR FROM Tanggal) as Tahun,
COUNT(case when choice='A' then 1 else NULL end) as A,
COUNT(case when choice='B' then 1 else NULL end) as B,
COUNT(case when choice='C' then 1 else NULL end) as C
FROM Table
GROUP BY extract(YEAR FROM Tanggal);