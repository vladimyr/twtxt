function reply(e) {
  e.preventDefault();

  var el = u("textarea#text")
  var text = document.getElementById("text");

  el.empty();
  el.text(u(e.target).data("reply"));
  el.scroll();

  text.focus();

  var size = el.text().length;

  text.setSelectionRange(size, size);
}

u(".reply").on("click", reply);

u("#post").on("click", function(e) {
  e.preventDefault();
  u("#post").html("<i class=\"icss-spinner icss-pulse\"></i>&nbsp;Posting...");
  u("#post").attr("disabled", true);
  u("#post").closest("form").first().submit();
});
