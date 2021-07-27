<template>
  <div
    :class="[{'rounded-full': rounded}, { indeterminate: indeterminate}]"
    class="w-full bg-gray-200 h-2 relative overflow-hidden"
  >
    <div
      :aria-valuenow="percentage"
      :class="[`bg-${color}-500`, {'absolute top-0': indeterminate}, {'rounded-full': rounded}]"
      :style="{width: `${percentage}%`}"
      aria-valuemax="100"
      aria-valuemin="0"
      class="h-full progressbar"
      role="progressbar"
    >
      <span class="flex items-center h-full">
        <slot></slot>
      </span>
    </div>
  </div>
</template>

<script>
  export default {
    inheritAttrs: false,
    props: {
      color: {
        type: String,
        default: "teal"
      },
      percentage: {
        type: Number,
        default: 0
      },
      rounded: {
        type: Boolean,
        default: true
      },
      indeterminate: Boolean
    }
  };
</script>
<style scoped>
  @keyframes progress-indeterminate {
    0% {
      width: 30%;
      left: -40%;
    }
    60% {
      left: 100%;
      width: 100%;
    }
    to {
      left: 100%;
      width: 0;
    }
  }

  .progressbar {
    transition: width 0.25s ease;
  }

  .indeterminate .progressbar {
    animation: progress-indeterminate 1.4s ease infinite;
  }
</style>
