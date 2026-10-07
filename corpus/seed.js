// FuzzForge starter seed
function seed(x) {
  let a = [x, 1, 2, 3];
  return a.map(v => v + 1).join(",");
}
seed(0);
