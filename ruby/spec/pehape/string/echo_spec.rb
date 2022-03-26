# frozen_string_literal: true

RSpec.describe Pehape do
  it "echo all type data" do
    expect { Pehape.echo("Pe ha pe") }.to output("Pe ha pe\n").to_stdout
    expect { Pehape.echo(nil) }.to output(nil).to_stdout
    expect { Pehape.echo(true) }.to output("true\n").to_stdout
    expect { Pehape.echo(1) }.to output("1\n").to_stdout
    expect { Pehape.echo(-3) }.to output("-3\n").to_stdout
    expect { Pehape.echo(["Pe", "Ha"]) }.to output("Pe\nHa\n").to_stdout
    expect { Pehape.echo({:Pe =>"Ha"}) }.to output("{:Pe=>\"Ha\"}\n").to_stdout
  end
end
