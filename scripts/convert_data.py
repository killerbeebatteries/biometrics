#!/usr/bin/env python3
#
# Grabbing the data from our source data file and importing to a postgres database.
#

import re
import datetime
import psycopg2

src_file = "../data/bio.txt"

formatted_data = []
timeperiod_comments = []

# read data into our array
with open(src_file, "r") as myfile:
    # remove blank lines
    # as per: https://stackoverflow.com/questions/4842057/easiest-way-to-ignore-blank-lines-when-reading-a-file-in-python
    src_data = (row.rstrip() for row in myfile)
    src_data = (row for row in src_data if row)

    # first record date
    date = "20170113"

    for row in src_data:
        # print row

        # if the row starts with a hash, then it is considered a general comment for that period of time.
        if row.startswith("#"):
            # use the date from the previous record.
            timeperiod_comments.append({"date": date, "comment": row})
            continue

        # the data from the start of the line to the first hash is the graph data.
        # everything after that is considered a comment (can be used as an annotation in a graph)

        graph_values = row.split("#")[0]

        # a bit of data cleanup
        # sometimes the graph values are seperated by one or more spaces. convert them into single-space
        graph_values = re.sub("\s+", " ", graph_values)

        # remove leading and trailing spaces
        graph_values = graph_values.strip()

        bp_and_weight_record = {}

        # sometimes there is no comment
        if len(row.split("#")) > 1:
            comment = row.split("#")[1]
        else:
            comment = ""

        bp_and_weight_record["comment"] = str(comment)
        # print "Test: " + graph_values
        # date going from YYmmdd format to YY-mm-dd format
        date = graph_values.split(" ")[0]
        date = datetime.datetime.strptime(date, "%Y%m%d")
        date = date.strftime("%Y-%m-%d")

        bp_sys = graph_values.split(" ")[1]
        bp_dsys = graph_values.split(" ")[2]
        bp_hb = graph_values.split(" ")[3]

        # sometimes there is no bp data
        haveBPData = True
        for bp_value in [bp_sys, bp_dsys, bp_hb]:
            if bp_value.upper() == "x".upper():
                haveBPData = False

        bp_and_weight_record["date"] = date

        if len(graph_values.split(" ")) > 7:
            bp_and_weight_record["time"] = graph_values.split(" ")[7]
        else:
            bp_and_weight_record["time"] = None

        if haveBPData:
            bp_and_weight_record["bp_sys"] = int(bp_sys)
            bp_and_weight_record["bp_dsys"] = int(bp_dsys)
            bp_and_weight_record["bp_hb"] = int(bp_hb)
        else:
            bp_and_weight_record["bp_sys"] = None
            bp_and_weight_record["bp_dsys"] = None
            bp_and_weight_record["bp_hb"] = None

        # sometimes we have no weight data
        if len(graph_values.split(" ")) > 5:
            weight_total = graph_values.split(" ")[4]
            weight_fat = graph_values.split(" ")[5]
            weight_muscle = graph_values.split(" ")[6]

            # sometimes we're missing only some of the weight data
            weight_total = float(weight_total) if weight_total.upper() != "X" else None
            weight_fat = float(weight_fat) if weight_fat.upper() != "X" else None
            weight_muscle = (
                float(weight_muscle) if weight_muscle.upper() != "X" else None
            )

            bp_and_weight_record["weight_total"] = weight_total
            bp_and_weight_record["weight_fat"] = weight_fat
            bp_and_weight_record["weight_muscle"] = weight_muscle
        else:
            bp_and_weight_record["weight_total"] = None
            bp_and_weight_record["weight_fat"] = None
            bp_and_weight_record["weight_muscle"] = None

        formatted_data.append(bp_and_weight_record)

# print(formatted_data)
# print(timeperiod_comments)


# connection establishment
conn = psycopg2.connect(
    database="biometrics",
    user="biometrics_user",
    password="somethingaboutacatonthemat",
    host="localhost",
    port="5432",
)

cur = conn.cursor()
cur.executemany(
    """INSERT INTO bp_and_weight(date, time, sys, dia, bp, weight_total, weight_fat, weight_muscle, comment) 
                VALUES (
                    %(date)s,
                    %(time)s,
                    %(bp_sys)s,
                    %(bp_dsys)s,
                    %(bp_hb)s,
                    %(weight_total)s,
                    %(weight_fat)s,
                    %(weight_muscle)s,
                    %(comment)s
                )
                """,
    formatted_data,
)

cur.executemany(
    """INSERT INTO bp_and_weight_comments(date, comment)
                VALUES (
                    %(date)s,
                    %(comment)s
                )
                """,
    timeperiod_comments,
)


sql = """select * from bp_and_weight;"""
cur.execute(sql)

for i in cur.fetchall():
    print(i)

sql = """select * from bp_and_weight_comments;"""
cur.execute(sql)

for i in cur.fetchall():
    print(i)

conn.commit()
# conn.rollback()
conn.close()
