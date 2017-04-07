package cashier

type Worker interface {}

type worker struct {
  pckg string
  function string
}

func NewWorker(pckg, function string) Worker {
  return &worker{
    pckg: pckg,
    function: function,
  }
}
