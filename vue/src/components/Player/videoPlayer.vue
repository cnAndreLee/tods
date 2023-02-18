<template>
    <vue3VideoPlay
        id="player"
        v-bind="options"
        poster="https://cdn.jsdelivr.net/gh/xdlumia/files/video-play/ironMan.jpg" >
    </vue3VideoPlay>
</template>

<script setup>
import { ref ,reactive ,computed, onBeforeUpdate } from 'vue';
import { useStore } from 'vuex'
import "vue3-video-play/dist/style.css";
import vue3VideoPlay from "vue3-video-play";
import config from '../../config';

const store = useStore()

const selectedFileUrl = computed(() => {
  const selectedFile = store.state.fileModule.selectedFile
  const url = config.backend + "/api/files/" + selectedFile.id + "." + selectedFile.suffix
  return url
})

const options = reactive({
  width: "100%", //播放器高度
  height: "650px", //播放器高度
  color: "#409eff", //主题色
  title: "", //视频名称
  src: selectedFileUrl, //视频源
  muted: false, //静音
  webFullScreen: false,
  speedRate: ["0.75", "1.0", "1.25", "1.5", "2.0"], //播放倍速
  autoPlay: false, //自动播放
  loop: false, //循环播放
  mirror: false, //镜像画面
  ligthOff: false, //关灯模式
  volume: 0.3, //默认音量大小
  control: true, //是否显示控制
  controlBtns: [
    "audioTrack",
    "quality",
    "speedRate",
    "volume",
    "setting",
    "pip",
    "pageFullScreen",
    "fullScreen",
  ], //显示所有按钮,
});

onBeforeUpdate(()=>{
  // 暂停视频
})

</script>