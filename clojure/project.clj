(defproject com.teknologiumum/pehape-clj "0.0.1"
  :description "PHP stdlib global functions implemented in Clojure"
  :url "https://github.com/teknologi-umum/pehape"
  :plugins [[lein-cloverage "1.2.2"]]
  :dependencies [[org.clojure/clojure "1.10.3"]]
  :profiles {:dev {:dependencies [[cider/cider-nrepl "0.28.3"]]}}
  :target-path "target/%s"
  :main pehape.core
  :aot [pehape.core])
