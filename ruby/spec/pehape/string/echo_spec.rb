# frozen_string_literal: true

RSpec.describe Pehape do
  it "echo all type data" do
    expect(Pehape.echo("Pe ha pe")).to equal(nil)
    expect(Pehape.echo(nil)).to equal(nil)
    expect(Pehape.echo(true)).to equal(nil)
    expect(Pehape.echo(1)).to equal(nil)
    expect(Pehape.echo(-3)).to equal(nil)
    expect(Pehape.echo(["Pe", "Ha"])).to equal(nil)
    expect(Pehape.echo({:Pe =>"Ha"})).to equal(nil)
  end
end
