#!/usr/bin/env python
#
# Grabbing the data from our source data file and converting to csv for use with graphing tools.
#

import re
import datetime
import plotly.graph_objects as go

src_file = "data/pulse.txt"

bp_data = []
weight_data = []
timeperiod_comments = []

# read data into our array
with open(src_file, 'r') as myfile:

    # remove blank lines
    # as per: https://stackoverflow.com/questions/4842057/easiest-way-to-ignore-blank-lines-when-reading-a-file-in-python
    src_data = (row.rstrip() for row in myfile)
    src_data = (row for row in src_data if row)


    for row in src_data:
        #print row

        # if the row starts with a hash, then it is considered a general comment for that period of time.
        if row.startswith("#"):
            # use the date from the previous record.
            timeperiod_comments.append({ "date": date, "comment": row })
            continue

        # the data from the start of the line to the first hash is the graph data.
        # everything after that is considered a comment (can be used as an annotation in a graph)

        graph_values = row.split("#")[0]

        # a bit of data cleanup
        # sometimes the graph values are seperated by one or more spaces. convert them into single-space
        graph_values = re.sub('\s+', ' ', graph_values)

        # remove leading and trailing spaces
        graph_values = graph_values.strip()

        # sometimes there is no comment
        if len(row.split("#")) > 1:
            comment = row.split("#")[1]
        else:
            comment = ""

        #print "Test: " + graph_values
        # date going from YYmmdd format to YY-mm-dd format
        date    = graph_values.split(" ")[0]
        date    = datetime.datetime.strptime(date, '%Y%m%d')
        date    = date.strftime('%Y-%m-%d')

        bp_sys  = graph_values.split(" ")[1]
        bp_dsys = graph_values.split(" ")[2]
        bp_hb =   graph_values.split(" ")[3]

        # sometimes there is no bp data
        haveBPData = True
        for bp_value in [ bp_sys, bp_dsys, bp_hb ]:
            if bp_value.upper() == "x".upper():
                haveBPData = False

        if haveBPData:
            bp_data.append({
                "date": date,
                "bp_sys": int(bp_sys),
                "bp_dsys": int(bp_dsys),
                "bp_hb": int(bp_hb)
                })


        # sometimes we have no weight data
        if len(graph_values.split(" ")) > 5:
            weight_total  = graph_values.split(" ")[4]
            weight_fat    = graph_values.split(" ")[5]
            weight_muscle = graph_values.split(" ")[6]

            weight_data.append({
                "date": date,
                "weight_total": float(weight_total),
                "weight_fat": float(weight_fat),
                "weight_muscle": float(weight_muscle)
            })

#print bp_data
#print weight_data

        # graph our data
