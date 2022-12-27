<template>
    <divTable :columnObj="columnObj" :tableData="tableData" />
</template>

<script setup>
import divTable from '../../components/divTable/index.vue'

const columnObj = {
    headers: [
        {
            prop: "rank",
            label: "排名",
            width: "5%",
        },
        {
            prop: "name",
            label: "名称",
            width: "15%",
        },
        {
            prop: "wcbfb",
            label: "完成百分比",
            width: "20%",
        },
        {
            prop: "wcrwl",
            label: "完成任务量",
            width: "20%",
        },
        {
            prop: "mbrwl",
            label: "目标任务量",
            width: "20%",
        },
        {
            prop: "syhkl",
            label: "剩余回款量",
            width: "20%",
        },
    ],
    rowStyle: ({ row }) => {
        let style = {}
        switch (row.rank) {
            case 1:
                style = {
                    backgroundColor: 'green',
                    color: '#000',
                }
                break;
            default:
                style = {
                    backgroundColor: 'red',
                    color: '#000'
                }
        }
        return style;
    },
}

const tableData = [
    {
        rank: 1,
        name: '北京一区',
        wcbfb: 50,
        wcrwl: 10000,
        mbrwl: 20000,
        syhkl: 9999
    },
    {
        rank: 2,
        name: '北京二区',
        wcbfb: 10,
        wcrwl: 5000,
        mbrwl: 50000,
        syhkl: 8888
    },
]
</script>