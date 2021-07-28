<template>
  <div>
    <h1 class="text-6xl tracking-wide	font-extrabold text-gray-900 mt-5">
      {{ quizScore }}/100</h1>
    <div @click="copyLink" class="mt-3 font-bold cursor-pointer">
      TEST ID: {{ $route.params.id }}
    </div>
    <div class="flex flex-nowrap justify-center mt-5">
      <div @click="tryAgain" class="mx-4 underline cursor-pointer">
        try again
      </div>
      <div @click="newQuiz" class="mx-4 underline cursor-pointer">
        new quiz
      </div>
    </div>
    <div class="mt-5">
      <div @click="showAnswer=!showAnswer" class="cursor-pointer">
        <p class="underline" v-if="!showAnswer">show correct answer</p>
        <p class="underline" v-else>hide correct answer</p>
      </div>
      <div v-show="showAnswer">
        <div class="flex flex-wrap px-4 mt-2">
          <progress-bar :percentage="currentProgress" class="mx-2 mb-2 h-5" color="gray">
            <span class="text-xs text-white w-full flex justify-end pr-2">{{currentProgress | mathCeil }}%</span>
          </progress-bar>
        </div>
        <div>click id/name to play full/quiz audio</div>
        <table class="border-collapse border border-gray-900 w-full mt-2">
          <tr v-for="(data,index) in quizInfo.pieces">
            <td @click="playPiece(data.path)" class="w-1/12  border border-gray-900 cursor-pointer">
              {{ index+1 }}
            </td>
            <td :class="failedQuiz.includes(index+1)?'bg-red-300':'bg-green-300'" @click="playFull(data.full)"
                class="w-11/12 py-2 border border-gray-500 cursor-pointer">
              {{ data.title }}
            </td>
          </tr>
        </table>
      </div>
    </div>
  </div>
</template>

<script>
  import ProgressBar from "../../components/ProgressBar";

  var audio = new Audio();

  export default {
    name: "Score",
    props: ['quizScore', 'quizInfo', 'failedQuiz'],
    components: {
      ProgressBar
    },
    data() {
      return {
        showAnswer: false,
        shareLink: window.location.protocol + "//" + window.location.host + "/start/" + this.$route.params.id,
        currentProgress: 0,
      }
    },
    methods: {
      copyLink() {
        window.prompt("Copy to clipboard: Ctrl+C, Enter", this.shareLink);
      },
      updateProgress() {
        var currentTime = audio.currentTime;
        var duration = audio.duration;
        if (isNaN(currentTime) || isNaN(duration)) {
          return
        }
        this.currentProgress = currentTime * 100 / duration
      },
      playMusic(url) {
        audio.pause();
        audio.removeEventListener("timeupdate", this.updateProgress);
        audio = new Audio(url);
        audio.play();
        audio.addEventListener("timeupdate", this.updateProgress);
      },
      playPiece(path) {
        this.playMusic(sessionStorage.getItem(path))
      },
      playFull(path) {
        this.playMusic(this.quizInfo.base + "/" + path)
      },
      tryAgain() {
        this.$router.push({
          name: 'Question', params: {
            quizInfo: this.quizInfo
          }
        })
      },
      newQuiz() {
        this.$router.push({name: 'Home'})
      }
    },
    mounted() {
      if (this.quizScore == null || this.quizInfo == null) {
        this.$router.push({name: 'Home'})
      }
    },
    filters: {
      mathCeil(i) {
        return Math.ceil(i)
      }
    }
  }
</script>

<style scoped>

</style>
