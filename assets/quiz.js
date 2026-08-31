/**
 * Reusable quiz widget for gips lessons.
 * All options in a quiz should be equal word count to avoid answer leakage.
 */
(function () {
  function initQuiz(container) {
    const correctIndex = parseInt(container.dataset.correct, 10);
    const feedbackOk = container.dataset.feedbackOk || "Correct.";
    const feedbackErr = container.dataset.feedbackErr || "Not quite — read the explanation above and try again.";
    const options = container.querySelectorAll(".quiz-option");
    const feedback = container.querySelector(".quiz-feedback");

    options.forEach(function (btn, index) {
      btn.addEventListener("click", function () {
        if (container.dataset.answered === "true") return;
        container.dataset.answered = "true";

        options.forEach(function (b) { b.disabled = true; });

        if (index === correctIndex) {
          btn.classList.add("correct");
          feedback.textContent = feedbackOk;
          feedback.className = "quiz-feedback ok";
        } else {
          btn.classList.add("incorrect");
          options[correctIndex].classList.add("correct");
          feedback.textContent = feedbackErr;
          feedback.className = "quiz-feedback err";
        }
      });
    });
  }

  document.addEventListener("DOMContentLoaded", function () {
    document.querySelectorAll(".quiz").forEach(initQuiz);
  });
})();
