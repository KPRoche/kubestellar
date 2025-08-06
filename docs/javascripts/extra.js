<script>
    document.addEventListener("DOMContentLoaded", function () {
      const banner = document.getElementById("ks-vers-warning");
      let lastScroll = window.scrollY;

      function handleScroll() {
        const currentScroll = window.scrollY;
        const isMobile = window.innerWidth <= 768;

        if (!isMobile) {
          banner.style.transform = "none";
          return;
        }

        if (currentScroll > lastScroll) {
          //banner.style.transform = "translateY(-120%)";
          banner.style.display = "none";
        } else {
          banner.style.display = "block";
          //banner.style.transform = "translateY(0)";
        }
        lastScroll = currentScroll;
      }

      banner.style.transition = "transform 0.3s ease";
      window.addEventListener("scroll", handleScroll);
    });
  </script>