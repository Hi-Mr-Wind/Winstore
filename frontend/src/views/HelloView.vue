<script lang="ts" setup>
import {ref,reactive} from 'vue'
import {EventsOn} from "../../wailsjs/runtime";
import {DownloadFile} from "../../wailsjs/go/apps/App";
interface DownloadData {
  percentage: Number,
  downloaded: String,
  total: String,
  speed: String

}
interface JsonData {
  code: Number,
  message: string,
  data: DownloadData
}

const colors = [
  { color: '#f56c6c', percentage: 20 },
  { color: '#e6a23c', percentage: 40 },
  { color: '#5cb87a', percentage: 60 },
  { color: '#6f7ad3', percentage: 80 },
  { color: '#1989fa', percentage: 100 },
]

const downloadData = ref<DownloadData | null>(null)
const responseData = ref<JsonData | null>(null);
EventsOn("download", (data:string) => {
  const parsedData: JsonData = JSON.parse(data)
  console.log("download 事件触发：", parsedData.data)
  responseData.value = parsedData
  console.log("动态数据",downloadData.value)
    downloadData.value = {
      percentage: parsedData.data.percentage,
      downloaded: parsedData.data.downloaded,
      total: parsedData.data.total,
      speed: parsedData.data.speed
    }
  // downloadData.percentage = data.percentage
  // downloadData.downloaded = data.downloaded
  // downloadData.total = data.total
  // downloadData.speed = data.speed
  console.log("赋值动态数据",downloadData.value)
})

const down = () => {
  DownloadFile("https://dldir1v6.qq.com/weixin/Universal/Windows/WeChatWin.exe", "E:\\codes\\winstore","weiChat.exe").then()
}
</script>

<template>
  <div>
    <el-button @click="down">下载</el-button>
  </div>
  <div>
<!--    <div><el-progress :percentage="downloadData?.percentage"/></div>-->
    <div> <el-progress :text-inside="true" :stroke-width="26" :percentage="downloadData?.percentage" :color="colors" /></div>
    <h1>已下载：{{ downloadData?.downloaded }}</h1>
    <h1>总大小：{{ downloadData?.total }}</h1>
    <h1>下载速度：{{ downloadData?.speed}}</h1>
  </div>
</template>

<style scoped>

</style>