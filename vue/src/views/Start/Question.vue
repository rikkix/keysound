<template>
  <div>
    <div class="flex flex-wrap px-4">
      <progress-bar :percentage="totalProgress" class="mx-2 mb-2 h-5" color="gray">
        <span class="text-xs text-white w-full flex justify-end pr-2">
          {{currentQuizID}}/{{totalQuizID}}</span>
      </progress-bar>
    </div>
    <div class="flex flex-wrap px-4">
      <progress-bar :percentage="currentProgress" class="mx-2 mb-2 h-5" color="gray">
        <span class="text-xs text-white w-full flex justify-end pr-2">{{currentProgress}}%</span>
      </progress-bar>
    </div>
    <quiz :currentQuiz="currentQuiz" :each="each" :nextQuiz="nextQuiz" v-if="!isEnd"></quiz>
    <button :disabled="isEnd" @click="nextQuiz(0)" class="text-xs font-thin underline">skip</button>
  </div>
</template>

<script>
  import ProgressBar from "../../components/ProgressBar";
  import Quiz from '../../components/Quiz.vue'


  var audio;


  export default {
    name: "Question",
    props: ['quizInfo'],
    components: {
      Quiz,
      ProgressBar
    },
    data() {
      return {
        currentQuiz: this.quizInfo.pieces[0],
        currentQuizID: 0,
        totalQuizID: this.quizInfo.pieces.length,
        totalProgress: 0,
        currentProgress: 0,
        isEnd: false,
        each: 100 / this.quizInfo.pieces.length,
        score: 0,
        failedQuiz: []
      }
    },
    methods: {
      updateProgress() {
        var currentTime = audio.currentTime;
        var duration = audio.duration;
        if (isNaN(currentTime) || isNaN(duration)) {
          return
        }
        this.currentProgress = Math.ceil(currentTime * 100 / duration)
      },
      nextQuiz(score) {
        if (score <= 0) {
          this.failedQuiz.push(this.currentQuizID)
        }
        this.score += score;
        audio.pause();
        audio.removeEventListener("timeupdate", this.updateProgress);
        if (this.currentQuizID === this.quizInfo.pieces.length) {
          this.isEnd = true;
          console.log(this.quizInfo);
          this.$router.push({
            name: 'Score',
            params: {
              quizInfo: this.quizInfo,
              quizScore: this.score,
              failedQuiz: this.failedQuiz
            }
          });
          return
        }
        this.currentQuiz = this.quizInfo.pieces[this.currentQuizID];
        audio = new Audio(sessionStorage.getItem(this.currentQuiz.path));
        audio.play();
        audio.addEventListener("timeupdate", this.updateProgress);
        this.totalProgress += this.each;
        this.currentQuizID++
      }
    },
    destroyed() {
      audio.pause()
    },
    mounted() {
      if (this.quizInfo == null) {
        this.$router.push({name: 'Home'})
      }
      audio = new Audio(sessionStorage.getItem(this.currentQuiz.path));
      audio.play();
      audio.addEventListener("timeupdate", this.updateProgress);
      this.totalProgress += this.each;
      this.currentQuizID++
    }
  }
</script>

<style scoped>

</style>
