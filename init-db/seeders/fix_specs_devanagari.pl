#!/usr/bin/perl
use strict;
use warnings;

my %devanagari_to_ascii = (
    '०' => '0', '१' => '1', '२' => '2', '३' => '3', '४' => '4',
    '५' => '5', '६' => '6', '७' => '7', '८' => '8', '९' => '9',
);

my $in_specs_map = 0;
my @lines;

while (<>) {
    my $line = $_;
    
    # Detect if we're in the specs map
    if ($line =~ /^\s*specs\s*:=\s*map\[string\]string\{/) {
        $in_specs_map = 1;
    }
    
    # Detect end of specs map
    if ($in_specs_map && $line =~ /^\s*\}/) {
        $in_specs_map = 0;
    }
    
    # Convert Devanagari to ASCII only within specs map
    if ($in_specs_map && $line =~ /:/) {
        foreach my $dev (keys %devanagari_to_ascii) {
            $line =~ s/$dev/$devanagari_to_ascii{$dev}/g;
        }
    }
    
    push @lines, $line;
}

print @lines;
