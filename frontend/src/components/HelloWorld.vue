<script setup lang="ts">
import {reactive, ref} from 'vue'
import {DownloadFile, Greet, SelectPage} from '../../wailsjs/go/apps/App'
import {useAppStore} from '../js/app'
import {EventsOn} from '../../wailsjs/runtime'

const data = reactive({
  name: "",
  resultText: "请在下面输入您的姓名 👇",
})

function greet() {
  Greet(data.name).then(result => {
    data.resultText = result
    useAppStore().map.set("123", data.resultText)
  })
}

function selectPage() {
  SelectPage(1, 10, data.name).then(result => {
    data.resultText = result.data
    const v = useAppStore().map.get("123")
    console.log(v)
  })
}
// interface DownloadData {
//   percentage: Number,
//   downloaded: String,
//   total: String,
//   speed: String
// }
const downloadData = ref({})
EventsOn("download", (data) => {
  console.log("download 事件触发：", data)
  downloadData.value = data
  // downloadData.percentage = data.percentage
  // downloadData.downloaded = data.downloaded
  // downloadData.total = data.total
  // downloadData.speed = data.speed
})

const down = () => {
  // DownloadFile("https://dldir1v6.qq.com/weixin/Universal/Windows/WeChatWin.exe", "E:\\codes\\winstore\\weiChat.exe").then()
}

</script>

<template>
  <main>
    <div id="result" class="result">{{ data.resultText }}</div>
    <div id="input" class="input-box">
      <input id="name" v-model="data.name" autocomplete="off" class="input" type="text"/>
      <button class="btn" @click="greet">问候</button>
    </div>

    <div>
      <el-button class="el-but" type="primary" round @click="selectPage">查询</el-button>
    </div>

  </main>
</template>

<style scoped>

.el-but {
  margin: 10px;
}

.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
